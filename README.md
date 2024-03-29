# Axolotl

Axolotl is a complete cross-platform [Signal](https://www.signal.org) client, compatible with the Ubuntu Phone and more.
Unlike the desktop Signal client, **Axolotl is completely autonomous** and doesn't require you to have created an
account with the official Signal application.

It is built upon [presage](https://github.com/whisperfish/presage) and a VueJs frontend that runs in a [tauri](https://tauri.app/) or a qml WebEngineView container.

<p align="center">
  <kbd>
    <img src="https://raw.githubusercontent.com/nanu-c/axolotl/main/screenshot.png" alt="Screenshot of axolotl" width="300px"/>
  </kbd>
</p>

## Features

 * Phone registration
 * Direct messages
 * Group messages *mostly*
 * Photo, video, audio and contact attachments in both direct and group mode
 * Preview for photo and audio attachments
 * Storing conversations
 * Encrypted message store
 * Desktop client provisioning/syncing *partially*

 ## Broken:
 * Contact discovery


There are still bugs and UI/UX quirks.

## Installation

Axolotl can be installed through different means.

| Package | Maintainer | Comment  |
| ------- | ---------- | -------- |
| <a href='https://open-store.io/app/axolotl.nanuc'><img width='130' alt="Get it from the OpenStore" src="https://open-store.io/badges/en_US.png"></a> | nanu-c | For Ubuntu Touch |
| <a href='https://snapcraft.io/axolotl'><img width='130' alt="Get it from the Snap Store" src="https://snapcraft.io/static/images/badges/en/snap-store-black.svg"></a> | nanu-c | For Ubuntu desktop |
| <a href='https://flathub.org/apps/details/org.nanuc.Axolotl'><img width='130' alt='Download on Flathub' src='https://flathub.org/assets/badges/flathub-badge-en.png'/></a> | olof-nord | https://github.com/flathub/org.nanuc.Axolotl |
| <a href='https://github.com/nanu-c/axolotl/releases'><img alt="mobian version" src="https://img.shields.io/badge/axolotl-deb-%23A80030"></a> | nanu-c | https://github.com/nanu-c/axolotl/releases <br>(Debian package in the Assets section)|

## Building

To find out how to build from source and install yourself, please see below.

* with Clickable: see [here](docs/INSTALL.md#clickable).
* with Snap: see [here](docs/INSTALL.md#snap).
* with Flatpak: see [here](docs/INSTALL.md#flatpak).
* with AppImage: see [here](docs/INSTALL.md#appimage).
* for Debian: see [here](docs/INSTALL.md#debian).
* manually: see [here](docs/INSTALL.md#bare).

## Run flags

- `--mode`
  - `tauri`: Default GUI
  - `ubuntu-touch`: Ubuntu Touch GUI
  - `daemon`: Headless

## Contributing

* Please fill issues here on github https://github.com/nanu-c/axolotl/issues
* Help translate Axolotl to your language(s). For information how to translate, please see [TRANSLATE.md](docs/TRANSLATE.md).
* Contribute code by making PR's (pull requests)

If you contribute new strings, please:

- make them translatable
- avoid linebreaks within one tag, that will break extracting the strings for translation
- try to reduce formatting tags within translatable strings

Translation is done by using the `easygettext` module. Detailed instructions how strings are made translatable are given here: [https://www.npmjs.com/package/easygettext](https://www.npmjs.com/package/easygettext).

In short words, either use the `v-translate` keyword in the last tag enclosing the string or wrap your string in a `<translate>` tag as the last tag.
If you need to make strings in the script section translatable, do it like this `this.$gettext("string")`.

When adding new translatable strings with a PR, make sure to extract and update commands as instructed [here](docs/TRANSLATE.md). Then also commit the updated pot and po files containing the new strings.

examples:

- `<p v-translate>Translate me!</p>` instead of `<p>Translate me!</p>`
- `<p><strong v-translate>Translate me!</strong></p>` instead of `<p><strong>Translate me!</strong></p>`
- `<p v-translate>Translate me!</p><br/><p v-translate> Please...</p>` instead of `<p>Translate me! <br/> Please...</p>`
- `<div v-translate>Yes, I am translatable!</div>` instead of `<div>No, I am not translatable!</div>`
- `<div><translate>This is a free and open source Signal client written in golang and vuejs.</translate></div>`
- in \<script\> part: `this.cMTitle = this.$gettext("I am a translatable title!");`
