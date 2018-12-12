import axios from 'axios'

export const getAccessToken = function () {
  return localStorage.getItem('access_token')
}

export const request = (option) => {
  return new Promise((resolve, reject) => {
    axios(option).then(v => {
      if (v.data.code === 200) {
        resolve(v.data.data)
      } else {
        reject(v.data)
      }
    }).catch(e => {
      reject(e)
    })
  })
}

export const createApi = (url, data) => {
  let headers = {
    'Access-Token': getAccessToken(),
  }

  let param = {
    url,
    data,
    method: 'post',
    headers
  }

  return request(param)
}