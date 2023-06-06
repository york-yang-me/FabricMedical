import request from '@/utils/request'

// Get a list of character selections on the login screen
export function queryAccountList() {
  return request({
    url: '/queryAccountList',
    method: 'post'
  })
}

// Login
export function login(data) {
  return request({
    url: '/queryAccountList',
    method: 'post',
    data
  })
}
