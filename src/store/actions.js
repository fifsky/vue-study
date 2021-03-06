import co from 'co'
import { loginApi,userApi } from '../service'

export default {
  loginAction({commit},data){
    return co(function * () {
      let ret = yield loginApi(data)

      if(!ret.access_token) {
        throw "登录失败"
      }

      localStorage.setItem('access_token', ret.access_token)
      commit("setUserInfo",ret.user)
    })
  },
  currentUserAction({commit}){
    return co(function * () {
      let ret = yield userApi()
      commit("setUserInfo",ret)
    })
  },
}
