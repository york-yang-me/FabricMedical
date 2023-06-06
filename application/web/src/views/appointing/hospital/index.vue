<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>Account ID: {{ accountId }}</p>
      <p>username: {{ userName }}</p>
      <p>balance: ￥{{ balance }} yen</p>
    </el-alert>
    <div v-if="appointingList.length==0" style="text-align: center;">
      <el-alert
          title="can not get the information"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in appointingList" :key="index" :span="6" :offset="1">
        <el-card class="d-buy-card">
          <div slot="header" class="clearfix">
            <span>{{ val.appointing.appointingStatus }}</span>
            <el-button v-if="val.appointing.appointingStatus==='appointing'" style="float: right; padding: 3px 0" type="text" @click="updateAppointing(val,'done')">accept</el-button>
            <el-button v-if="val.appointing.appointingStatus==='appointing'" style="float: right; padding: 3px 6px" type="text" @click="updateAppointing(val,'cancelled')">cancel</el-button>
          </div>
          <div class="item">
            <el-tag>sequence ID: </el-tag>
            <span>{{ val.appointing.objectOfAppointing }}</span>
          </div>
          <div class="item">
            <el-tag type="success">patient ID: </el-tag>
            <span>{{ val.appointing.patient }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">hospital ID: </el-tag>
            <span>{{ val.appointing.hospital }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">create time: </el-tag>
            <span>{{ val.appointing.createTime }}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAppointingListByHospital, updateAppointing } from '@/api/appointing'

export default {
  name: 'AppointingHospital',
  data() {
    return {
      loading: true,
      appointingList: []
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
    queryAppointingListByHospital({ hospital: this.accountId }).then(response => {
      if (response !== null) {
        this.appointingList = response
      }
      this.loading = false
    }).catch(_ => {
      this.loading = false
    })
  },
  methods: {
    updateAppointing(item, type) {
      let tip = ''
      if (type === 'done') {
        tip = 'confirm accepting'
      } else {
        tip = 'cancel accepting'
      }
      this.$confirm('If you need' + tip + '?', 'tips', {
        confirmButtonText: 'confirm',
        cancelButtonText: 'cancel',
        type: 'success'
      }).then(() => {
        this.loading = true
        updateAppointing({
          patient: item.appointing.patient,
          hospital: item.appointing.hospital,
          objectOfAppointing: item.appointing.objectOfAppointing,
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

  .d-buy-card {
    width: 280px;
    height: 300px;
    margin: 18px;
  }
</style>
