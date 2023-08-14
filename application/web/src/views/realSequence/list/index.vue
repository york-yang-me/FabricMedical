<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>Account ID: {{ accountId }}</p>
      <p>username: {{ userName }}</p>
      <p>balance: ￥{{ balance }} yen</p>
      <p>When an assignment  or authorizing operation is initiated, the endorsement is true</p>
      <p>When the endorsement is false, the assignment or authorizing operation can only be initiated</p>
    </el-alert>
    <div v-if="realSequenceList.length==0" style="text-align: center;">
      <el-alert
        title="can not find the information"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in realSequenceList" :key="index" :span="6" :offset="1">
        <el-card class="realSequence-card">
          <div slot="header" class="clearfix">
            Endorse status:
            <span style="color: rgb(255, 0, 0);">{{ val.endorsement }}</span>
          </div>

          <div class="item">
            <el-tag>Account ID: </el-tag>
            <span>{{ val.realSequenceId }}</span>
          </div>
          <div class="item">
            <el-tag type="success">owner ID: </el-tag>
            <span>{{ val.owner }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">total length: </el-tag>
            <span>{{ val.totalLength }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">dna contents: </el-tag>
            <span>{{ val.dnaContents }}</span>
          </div>
          <el-rate v-if="roles[0] === 'admin'" />
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryRealSequenceList } from '@/api/realSequence'

export default {
  name: 'RealSequence',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('must be bigger than0'))
      } else {
        callback()
      }
    }
    return {
      loading: true,
      loadingDialog: false,
      realSequenceList: [],
      rules: {
        price: [
          { validator: checkArea, trigger: 'blur' }
        ],
      },
      accountList: [],
      valItem: {}
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'roles',
      'userName',
      'balance'
    ])
  },
  created() {
    if (this.roles[0] === 'admin') {
      queryRealSequenceList().then(response => {
        if (response !== null) {
          this.realSequenceList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    } else {
      queryRealSequenceList({ owner: this.accountId }).then(response => {
        if (response !== null) {
          this.realSequenceList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
  },
  methods: {
  }
}

</script>

<style>
  .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
  }

  .realSequence-card {
    width: 280px;
    height: 340px;
    margin: 18px;
  }
</style>
