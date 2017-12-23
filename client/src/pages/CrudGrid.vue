<template>
  <div>
    <v-layout>
      <v-flex md4="md4"></v-flex>
      <v-flex md8="md8">
        <v-form class="row jr" :inline="true" v-model="filters.model" v-if="filters.fields" :fields="filters.fields" @submit="doSearch"
          submitButtonText="Search" submitButtonIcon="search"></v-form>
      </v-flex>
    </v-layout>
    <v-card>
      <div>
        <v-btn class="green" router="router" fab="fab" absolute="absolute" top="top" right="right" dark="dark" :to="{name: 'create', params: {resource}}"
          v-if="options.create !== false">
          <v-icon>add</v-icon>
        </v-btn>
      </div>
      <v-data-table :headers="columns" :items="items" :total-items="pagination.totalItems" hide-actions="hide-actions" :pagination.sync="pagination"
        :loading="loading">
        <template slot="items" scope="props">
          <tr>
            <td :class="column.left? '': 'text-xs-right'" v-for="column in columns" v-html="getColumnData(props.item, column)"></td>
            <td v-if="actions !== false" width="160">
              <template v-for="(value, action) in actions">
                <v-btn v-if="['edit', 'delete'].indexOf(action) &lt; 0" router="router" primary="primary" fab="fab" small="small" dark="dark"
                  :to="{name: action, params: {resource,id:props.item.id}}">
                  <v-icon>{{action.icon ? action.icon : action}}</v-icon>
                </v-btn>
              </template>
              <v-btn v-if="options.edit !== false" dark="dark" primary="primary" fab="fab" small="small" :to="{name: 'edit', params: {resource,id:props.item.id}}">
                <v-icon>edit</v-icon>
              </v-btn>
              <!-- also you can try this: inline edit-->
              <!-- v-btn(v-if="options.edit !== false",dark,fab,success,small,@click.native.stop="showEdit(props.item)")-->
              <!--   v-icon() edit-->
              <v-btn v-if="options.delete !== false" fab="fab" small="small" @click="remove(props.item)">
                <v-icon>delete</v-icon>
              </v-btn>
            </td>
          </tr>
        </template>
      </v-data-table>
      <div class="jc">
        <v-pagination class="ma-3" v-model="pagination.page" :length="totalPages" circle="circle"></v-pagination>
      </div>
    </v-card>
    <v-dialog v-model="isShowEdit" width="70%">
      <v-card>
        <v-card-title>{{$t('Edit')}} #{{currentItem.id}}</v-card-title>
        <v-card-text>
          <v-form v-model="form.model" v-bind="form" method="patch" :action="resource+'/'+currentItem.id" @success="onSaveEdit"></v-form>
        </v-card-text>
        <v-card-actions actions="actions">
          <v-btn flat="flat" primary="primary" @click="isShowEdit = false">{{$t('Close')}}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
  const getDefaultData = () => {
    return {
      isShowEdit: false,
      form: {
        model: {},
        fields: {},
        rules: {},
        messages: {}
      },
      filters: {
        model: {},
        fields: null
      },
      loading: false,
      columns: [], // fetch grid setup params from server if `columns` is empty
      actions: {},
      options: {
        sort: 'id',
        create: false,
        update: true,
        delete: false
      },
      pagination: {
        page: 1,
        rowsPerPage: 10,
        sortBy: 'id',
        descending: true,
        totalItems: 0
      },
      currentItem: false,
      items: []
    }
  }
  export default {

    data: getDefaultData,

    watch: {
      '$i18n.locale' (val) {
        this.fetchGrid()
      },
      'pagination.page' (val) {
        this.fetchData()
      },
      'pagination.sortBy' (val) {
        this.fetchData()
      },
      'pagination.descending' (val) {
        this.fetchData()
      },
      '$route.params': 'refresh'
      // '$route.query': 'updateRoute'
    },
    methods: {
      fetchForm (item) {
        this.$http.get(`${this.resource}/form`, {
          params: {
            id: item.id
          }
        }).then(({
          data
        }) => {
          this.form = data
        })
      },
      onSaveEdit (data) {
        if (data.id) {
          this.isShowEdit = false
          this.fetchData()
        }
      },
      showEdit (item) {
        this.currentItem = item
        this.fetchForm(item)
        this.isShowEdit = true
      },
      preFetch () {
        let sort = this.pagination.sortBy
        if (this.pagination.descending) {
          sort = '-' + sort
        }
        this.$route.query.query = JSON.stringify(this.filters.model)
        this.$route.query.sort = sort
        this.$route.query.perPage = this.pagination.rowsPerPage
        this.$route.query.page = this.pagination.page
      },
      updateRoute () {
        this.$route.query.keepLayout = true
        console.log('update route')
        this.$router.go({
          path: this.$route.path,
          params: this.$route.params,
          query: this.$route.query
        })
      },
      doSearch () {
        this.pagination.page = 1
        this.fetchData()
      },
      filter (val, search) {
        return true
        // this.search = search
        // this.fetchData()
      },
      refresh () {
        Object.assign(this.$data, getDefaultData())
        this.fetchGrid().then(() => {})
        this.fetchData()
      },
      fetch () {
        if (this.columns.length <= 0) {
          // fetch grid params from server: e.g. /crud/users/grid
          this.fetchGrid()
        } else {
          // or define grid params manually
          this.fetchData()
        }
      },
      getColumnData (row, field) {
        // process fields like `type.name`
        let [l1, l2] = field.value.split('.')
        let value = row[l1]
        let tag = null
        if (l2) {
          value = row[l1] ? row[l1][l2] : null
        }
        if (field.type === 'image') {
          tag = 'img'
        }
        if (tag) {
          value = `<${tag} src="${value}" class="crud-grid-thumb" controls />`
        }
        return value
      },
      fetchGrid () {
        return new Promise((resolve, reject) => {
          this.$http.get(`${this.resource}/grid`).then(({
            data
          }) => {
            for (let k in data.columns) {
              data.columns[k].text = this.$t(data.columns[k].text)
            }
            this.columns = data.columns || {}
            this.actions = data.actions || {}
            this.filters = data.filters || {}
            this.options = data.options || {}

            if (this.options && this.options.sort) {
              let sortData = this.options.sort.split('-')
              let desc = sortData.length > 1
              let sortField = sortData.pop()

              // if (sortField.indexOf('.') < 0) {
              //   sortField = sortField
              // }
              this.pagination.sort = sortField
              this.pagination.descending = desc
            }
            resolve()
          })
        })
      },
      fetchData () {
        this.preFetch()
        this.$http.get(`${this.resource}`, {
          params: this.$route.query
        }).then(({
          data
        }) => {
          this.items = data.data
          this.pagination.totalItems = data.total
        })
      },
      remove (item) {
        // this.$alert('ok')
        console.log(`delete ${item.id}`)
      },
      next () {
        // console.log('next')
        this.pagination.page++
      }
    },
    computed: {
      resource () {
        return this.$route.params.resource
      },
      totalPages () {
        return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
      }
    },

    mounted () {

    },
    created () {
      this.$store.commit('setPageTitle', global.helper.i.titleize(global.helper.i.pluralize(this.resource)))
      this.fetchGrid().then(() => {})
      this.fetchData()
      // this.fetch()
    }

  }

</script>
