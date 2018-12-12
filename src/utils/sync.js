import co from 'co'
import { get } from './maping'
import { dialog } from './dialog'

const coProcess = (fn, options = {errHandle: null}) => {
  co(function* () {
    yield fn()
  }).catch(e => {
    console.log(e)
    if(get(e,"stack")){
      throw e
    }

    if (options.errHandle) {
      co(function * () {
        yield options.errHandle(e)
      })
    }else{
      let msg = get(e,"msg","未知错误"+JSON.stringify(e))
      dialog.alert(msg)
    }
  })
}

export const sync = (fn,options) => coProcess(fn,options)
