import request from '@/utils/request'

// Query assignment list (can search all, also by patient)
export function queryAppointingList(data) {
  return request({
    url: '/queryAppointingList',
    method: 'post',
    data
  })
}

// Search for hospitals by hospitalId (for patient search)
export function queryAppointingListByHospital(data) {
  return request({
    url: '/queryAppointingListByHospital',
    method: 'post',
    data
  })
}

// Update appingting status "done" or "cancelled"
export function updateAppointing(data) {
  return request({
    url: '/updateAppointing',
    method: 'post',
    data
  })
}

// Create Appointing
export function createAppointing(data) {
  return request({
    url: '/createAppointing',
    method: 'post',
    data
  })
}
