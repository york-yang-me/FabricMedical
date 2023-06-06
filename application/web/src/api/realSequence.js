import request from '@/utils/request'

// record dna sequence information(admin)
export function createRealSequence(data) {
  return request({
    url: '/createRealSequence',
    method: 'post',
    data
  })
}

// Get dna information
// (empty json{} can query all, specified proprietor can query the hospital information owned under the specified hospital)
export function queryRealSequenceList(data) {
  return request({
    url: '/queryRealSequenceList',
    method: 'post',
    data
  })
}
