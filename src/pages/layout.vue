<template>
  <Layout>
    <Header style="background: #fff;height: auto;line-height: normal">
      <div class="main">
        <Menu mode="horizontal" theme="light" active-name="1" class="left">
          <MenuItem name="1">
            <RouterLink to="/">
              <Icon type="ios-home"/>首页</RouterLink>
          </MenuItem>
          <MenuItem name="2">
            <RouterLink to="/user">
              <Icon type="ios-people"/>个人中心</RouterLink>
          </MenuItem>
          <MenuItem name="3">
            <a href="https://fifsky.com" target="_blank"
               rel="noopener">
              <Icon type="ios-link"/>
              VeryPay</a>
          </MenuItem>
        </Menu>
        <div class="right">
          <a class="action" href="#" @click="logOut">退出</a>
        </div>
      </div>
    </Header>
    <Content>
      <router-view></router-view>
    </Content>
  </Layout>
</template>
<script>
  import {mapActions} from 'vuex'

  export default {
    data () {
      return {}
    },
    methods: {
      ...mapActions(['currentUserAction']),
      logOut () {
        localStorage.removeItem('access_token')
        this.$router.push('/login')
      },
    },
    mounted () {
      this.currentUserAction()
    },
  }
</script>

<style lang="less">
  @pro-header-hover-bg: rgba(0, 0, 0, 0.025);
  @text-color: fade(#000, 65%);

  .ivu-menu-horizontal.ivu-menu-light:after{
    height: 0 !important;
  }

  .main {
    display: flex;
    width: 800px;
    margin: 0 auto;

    .left {
      flex: 1;
      display: flex;
      border-bottom: none;
    }

    .right {
      float: right;
      height: 100%;
      line-height: 60px;
      overflow: hidden;

      .action {
        cursor: pointer;
        padding: 0 12px;
        display: inline-block;
        transition: all 0.3s;
        height: 100%;

        > i {
          vertical-align: middle;
          color: @text-color;
        }

        &:hover {
          background: @pro-header-hover-bg;
        }

        &:global(.opened) {
          background: @pro-header-hover-bg;
        }
      }
    }
  }
</style>