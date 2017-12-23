<template lang="pug">
  <v-layout>
    <v-flex xs8="xs8">
      <v-form v-model="model" v-bind="$data" :method="method" :action="action" @success="onSuccess">
        <div class="my-4" slot="buttons">
          <v-btn class="grey" dark="dark" @click.native="$root.back()">
            <v-icon dark="dark" left="left">chevron_left </v-icon><span>{{$t('Back')}}</span>
          </v-btn>
          <v-btn primary="primary" dark="dark" type="submit">{{$t('Submit')}}
            <v-icon right="right" dark="dark">send</v-icon>
          </v-btn>
        </div>
      </v-form>
    </v-flex>
  </v-layout>
</template>

<script>
  export default {

    data () {
      return {
        model: {},
        fields: {},
        rules: {},
        messages: {}
      }
    },
    computed: {
      method () {
        return this.isEdit ? 'patch' : 'post'
      },
      action () {
        if (this.isEdit) {
          return `${this.resource}/${this.id}`
        } else {
          return `${this.resource}`
        }
      },
      isEdit () {
        return !!this.id
      },
      resource () {
        return this.$route.params.resource
      },
      id () {
        return this.$route.params.id
      }

    },
    watch: {
      '$route': 'fetch',
      'model': 'updateFields'
    },
    methods: {
      getFieldError (fieldName) {
        for (let k in this.errors) {
          let error = this.errors[k]
          if (error.field === fieldName) {
            return error.message
          }
        }
      },
      updateFields () {

      },
      fetch () {
        this.$http.get(`${this.resource}/form`, {
          params: {
            id: this.id
          }
        }).then(({
          data
        }) => {
          this.model = data.model
          this.fields = data.fields
          this.rules = data.rules
          this.messages = data.messages
        })
      },
      onSubmit () {

      },
      onSuccess (data) {
        this.$router.push({
          name: 'grid',
          params: {
            resource: this.resource
          }
        })
        if (data.id) {
          // this.$router.go(-1)
        }
      }
    },
    created () {
      let pageTitle = (this.isEdit ? 'Update' : 'Create') + ' ' + global.helper.i.titleize(global.helper.i.singularize(
        this.resource))
      this.$store.commit('setPageTitle', pageTitle)
    },
    mounted () {
      // this.$bus.showMessage('success', 'success')
      this.fetch()
    }
  }

</script>
