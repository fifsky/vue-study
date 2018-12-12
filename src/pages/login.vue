<template>
    <div style="width: 200px;margin: 200px auto">
        <h2>登录</h2>
        <Form ref="loginForm" :model="data" @keydown.enter.native="handleSubmit">
            <FormItem prop="userName">
                <Input v-model="data.user_name" placeholder="请输入用户名">
        <span slot="prepend">
          <Icon :size="16" type="ios-person"></Icon>
        </span>
                </Input>
            </FormItem>
            <FormItem prop="password">
                <Input type="password" v-model="data.password" placeholder="请输入密码">
        <span slot="prepend">
          <Icon :size="14" type="md-lock"></Icon>
        </span>
                </Input>
            </FormItem>
            <FormItem>
                <Button @click="handleSubmit" type="primary" long>登录</Button>
            </FormItem>
        </Form>
    </div>
</template>

<script>
  import {sync} from "../utils"
  import {mapActions} from 'vuex'

  export default {
    name: 'home',
    data() {
      return {
        data: {
          user_name: '',
          password: ''
        }
      }
    },
    methods: {
      ...mapActions(['loginAction']),
      handleSubmit() {
        let self = this
        sync(function* () {
          yield self.loginAction(self.data)
          self.$router.push("/")
        })
      }
    }
  }
</script>
