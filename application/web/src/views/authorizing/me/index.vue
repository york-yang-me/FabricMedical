<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>Account ID: {{ accountId }}</p>
      <p>username: {{ userName }}</p>
      <p>balance: ￥{{ balance }} yen</p>
    </el-alert>
    <div v-if="authorizingList.length==0" style="text-align: center;">
      <el-alert
        title="can not find the information"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in authorizingList" :key="index" :span="6" :offset="1">
        <el-card class="me-card">
          <div slot="header" class="clearfix">
            <span>{{ val.authorizingStatus }}</span>
            <el-button v-if="val.authorizingStatus!=='finish'&&val.authorizingStatus!=='expired'&&val.authorizingStatus!=='cancelled'" style="float: right; padding: 3px 0" type="text" @click="updateAuthorizing(val,'cancelled')">cancel</el-button>
            <el-button v-if="val.authorizingStatus==='In delivery'" style="float: right; padding: 3px 8px" type="text" @click="updateAuthorizing(val,'done')">confirm</el-button>
          </div>
          <div class="item">
            <el-tag>RealSequence ID: </el-tag>
            <span>{{ val.objectOfAuthorize }}</span>
          </div>
          <div class="item">
            <el-tag type="success">patient ID: </el-tag>
            <span>{{ val.patient }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">prize: </el-tag>
            <span>￥{{ val.price }} yen</span>
          </div>
          <div class="item">
            <el-tag type="warning">Expiration: </el-tag>
            <span>{{ val.salePeriod }} 天</span>
          </div>
          <div class="item">
            <el-tag type="info">create time: </el-tag>
            <span>{{ val.createTime }}</span>
          </div>
          <div class="item">
            <el-tag>hospital ID: </el-tag>
            <span v-if="val.hospital===''">waiting</span>
            <span>{{ val.hospital }}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAuthorizingList, updateAuthorizing } from '@/api/authorizing'
import RealSequence from "@/views/realSequence/list/index.vue";

export default {
  name: 'MeAuthorizing',
  components: {RealSequence},
  data() {
    return {
      loading: true,
      authorizingList: []
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'userName',
      'balance'
    ])
  },
  created() {
    queryAuthorizingList({ patient: this.accountId }).then(response => {
      if (response !== null) {
        this.authorizingList = response
      }
      this.loading = false
    }).catch(_ => {
      this.loading = false
    })
  },
  methods: {
    updateAuthorizing(item, type) {
      let tip = ''
      if (type === 'done') {
        tip = 'confirm'
      } else {
        tip = 'cancel'
      }
      this.$confirm('If you need' + tip + '?', 'tip', {
        confirmButtonText: 'confirm',
        cancelButtonText: 'cancel',
        type: 'success'
      }).then(() => {
        this.loading = true
        updateAuthorizing({
          hospital: item.hospital,
          objectOfAuthorize: item.objectOfAuthorize,
          patient: item.patient,
          status: type
        }).then(response => {
          this.loading = false
          if (response !== null) {
            this.$message({
              type: 'success',
              message: tip + 'operation success!'
            })
          } else {
            this.$message({
              type: 'error',
              message: tip + 'operation failed!'
            })
          }
          setTimeout(() => {
            window.location.reload()
          }, 1000)
        }).catch(_ => {
          this.loading = false
        })
      }).catch(() => {
        this.loading = false
        this.$message({
          type: 'info',
          message: 'cancelled' + tip
        })
      })
    }
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
  .tag {
    float: left;
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

  .me-card {
    width: 280px;
    height: 380px;
    margin: 18px;
  }
</style>
