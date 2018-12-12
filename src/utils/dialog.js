export const dialog = {
  alert(msg,fn = null,title="提示"){
    window.alert(msg);
    fn && fn()
  },
  confirm(msg,fn=null,title="提示"){
    if(window.confirm(msg)){
      fn && fn(true)
    }else{
      fn && fn(false)
    }
  },
}
