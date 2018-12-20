<template>
  <div>
    <Card style="width:800px;margin: 20px auto" dis-hover>
      <p slot="title">个人中心</p>
      <p>用户：{{userInfo.name}}</p>
      <p>邮箱：{{userInfo.email}}</p>
      <p>注册时间：{{userInfo.created_at}}</p>
    </Card>
  </div>
</template>
<script>
  import {mapState} from 'vuex'
  export default {
    computed: {
      ...mapState(['userInfo']),
    },
  }
</script>