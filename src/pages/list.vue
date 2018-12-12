<template>
    <div>
        <Card style="width:780px;margin: 20px auto">
            <p slot="title">每日心情</p>
            <a href="#" slot="extra" @click.prevent="logOut">
                <Icon type="ios-loop-strong"></Icon>
                退出登录
            </a>
            <div style="padding: 0 0 10px 0;text-align: left">
                <Button type="primary" @click="add_modal=true">发表</Button>
            </div>
            <Table :loading="loading" stripe border :columns="columns" :data="data"></Table>
            <div style="margin-top: 10px">
                <Page :total="pager.total" :page-size="pager.num" show-total @on-change="getList"/>
            </div>
        </Card>
        <Modal
                v-model="add_modal"
                title="今天你😊吗？"
                @on-ok="add">
            <Form :model="formItem">
                <FormItem style="margin-bottom: 0">
                    <Input v-model="formItem.content" type="textarea" autofocus :autosize="{minRows: 4,maxRows: 10}" placeholder="写下你的心情......"></Input>
                </FormItem>
            </Form>

        </Modal>
        <Modal footer-hide
               v-model="view_modal"
                :title="get(view_data,'user.name') +'-'+ get(view_data,'created_at')"
                @on-ok="add">
        <div>
            <p>{{view_data.content}}</p>
        </div>
        </Modal>
    </div>
</template>
<script>
  import {get, sync} from "../utils";
  import {addApi, deleteApi, listApi} from '../api'

  export default {
    data() {
      return {
        "view_modal":false,
        "add_modal": false,
        "view_data":{},
        "loading": false,
        "formItem":{
          content:""
        },
        "columns": [
          {
            title: '作者',
            key: 'user',
            width: 100,
            render: (h, params) => {
              return h(
                'span', get(params.row, 'user.name')
              )
            }
          },
          {
            title: '心情',
            key: 'content'
          },
          {
            title: '日期',
            key: 'created_at',
            width: 150
          },
          {
            title: '操作',
            key: 'action',
            fixed: 'right',
            width: 120,
            render: (h, params) => {
              return h('div', [
                h('a', {
                  props: {},
                  style: {
                    marginRight: '5px'
                  },
                  on: {
                    click: () => {
                      this.view_modal = true
                      this.view_data = params.row
                    }
                  }
                }, '查看'),
                h('a', {
                  props: {},
                  style: {
                    marginRight: '5px'
                  },
                  on: {
                    click: () => {
                      this.$Modal.confirm({
                        title: '确认删除?',
                        onOk: () => {
                          this.remove(params)
                        }
                      });
                    }
                  }
                }, '删除')
              ]);
            }
          }
        ],
        "data": [],
        "pager": {
          "current": 1,
          "num": 10,
          "total": 0
        }
      }
    },
    methods: {
      get,
      logOut() {
        localStorage.removeItem("access_token")
        this.$router.push("/login")
      },
      getList(page) {
        let self = this
        sync(function* () {
          self.loading = true
          let ret = yield listApi({page: page})
          self.loading = false
          self.data = ret.list
          self.pager = ret.page
        })
      },
      remove(params) {
        let self = this
        sync(function* () {
          self.loading = true
          yield deleteApi({id: params.row.id})
          self.loading = false
          self.data.splice(params.index, 1);
          self.pager.total--
        })
      },
      add() {
        let self = this
        sync(function* () {
          self.loading = true
          let ret = yield addApi(self.formItem)
          self.formItem.content = ""
          self.loading = false
          self.data.unshift(ret)
          self.pager.total++
        })
      },
    },
    mounted() {
      this.getList(1);
    }
  }
</script>