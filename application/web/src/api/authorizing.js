import request from '@/utils/request'

// Query authorizing list(Query all or query by patient)
export function queryAuthorizingList(data) {
  return request({
    url: '/queryAuthorizingList',
    method: 'post',
    data
  })
}

// According to Hospital and patient (Hospital AccountId) query involved information
export function queryAuthorizingListByBuyer(data) {
  return request({
    url: '/queryAuthorizingListByBuyer',
    method: 'post',
    data
  })
}

// Hospital access
export function createAuthorizingByBuy(data) {
  return request({
    url: '/createAuthorizingByBuy',
    method: 'post',
    data
  })
}

// Update authorizing status "done" or "cancelled" When authorizing, the hospital want to cancel to store this, hospital is null
export function updateAuthorizing(data) {
  return request({
    url: '/updateAuthorizing',
    method: 'post',
    data
  })
}

// Create authorizing
export function createAuthorizing(data) {
  return request({
    url: '/createAuthorizing',
    method: 'post',
    data
  })
}
