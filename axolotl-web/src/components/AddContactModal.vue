<template>
  <div class="modal" tabindex="-1" role="dialog">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">
            <span v-translate>Add</span>
            <div v-if="name !== ''">{{ name }}</div>
            <div v-else v-translate>Contact</div>
          </h5>
          <button type="button" class="close" @click="$emit('close')">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <span><strong v-translate>Only add numbers that you are sure are registered Signal
            accounts!</strong></span>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label v-translate for="nameInput">Name</label>
            <input
              id="nameInput"
              v-model="name"
              type="text"
              class="form-control"
            >
          </div>
          <div class="form-group">
            <label v-translate for="phoneInput">Phone</label>
            <input
              id="phoneInput"
              v-model="phone"
              type="text"
              class="form-control"
              placeholder="+44..."
            >
          </div>
        </div>
        <div class="modal-footer">
          <button
            v-translate
            type="button"
            class="btn btn-primary"
            @click="$emit('add', { name: name, phone: phone })"
          >
            Add
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "AddContactModal",
  data() {
    return {
      phone: "",
      name: "",
    };
  },
};
</script>
<style scoped>
.modal {
  display: block;
  border: none;
}
.modal-content {
  border-radius: 0px;
}
.modal-header {
  border-bottom: none;
}
.modal-title {
  display: flex;
}
.modal-title > div {
  margin-left: 10px;
}
.modal-footer {
  border-top: 0px;
}
</style>
