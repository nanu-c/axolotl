package store

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/nanu-c/axolotl/app/config"
	"github.com/signal-golang/textsecure"
	textsecureContacts "github.com/signal-golang/textsecure/contacts"

	yaml "gopkg.in/yaml.v2"
)

type Contacts struct {
	Contacts []textsecureContacts.Contact
	Len      int
}

var ContactsModel *Contacts = &Contacts{} // TODO

func (c *Contacts) GetContact(i int) textsecureContacts.Contact {
	if i == -1 {
		return textsecureContacts.Contact{}
	}
	return c.Contacts[i]
}
func GetContactForTel(tel string) *textsecureContacts.Contact {
	for _, c := range ContactsModel.Contacts {
		if c.Tel == tel {
			return &c
		}
	}
	return nil
}
func GetContactForUUID(uuid string) *textsecureContacts.Contact {
	for _, c := range ContactsModel.Contacts {
		if c.UUID == uuid {
			return &c
		}
	}
	return nil
}
func RefreshContacts() error {
	registeredContactsFile := config.GetRegisteredContactsFile()
	contacts, err := textsecure.GetRegisteredContacts()
	if err != nil {
		log.Errorln("[axolotl] RefreshContacts", err)
		// when refresh fails, we load the old contacts
		contacts, _ = readRegisteredContacts(registeredContactsFile)
		return err
	} else {
		writeRegisteredContacts(registeredContactsFile, contacts)
	}
	log.Debugln("[axolotl] Refresh contacts count: ", len(contacts))
	ContactsModel.Contacts = contacts
	ContactsModel.Len = len(contacts)
	SessionsModel.UpdateSessionNames()
	if err != nil {
		return err
	}
	return nil
}

type yamlContacts struct {
	Contacts []textsecureContacts.Contact
}

func writeRegisteredContacts(filename string, contacts []textsecureContacts.Contact) error {
	c := &yamlContacts{contacts}
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0600)
}
func readRegisteredContacts(fileName string) ([]textsecureContacts.Contact, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	contacts := &yamlContacts{}
	err = yaml.Unmarshal(b, contacts)
	if err != nil {
		return nil, err
	}
	return contacts.Contacts, nil
}

func TelToName(tel string) string {
	if g, ok := Groups[tel]; ok {
		return g.Name
	}
	for _, c := range ContactsModel.Contacts {
		if c.Tel == tel {
			return c.Name
		}
	}
	config := &config.Config{} // TODO - wire through config
	if tel == config.GetMyNumber() {
		return "Me"
	}
	return tel
}
