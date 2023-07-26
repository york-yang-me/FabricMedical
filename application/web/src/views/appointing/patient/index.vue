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
        <el-card class="d-me-card">
          <div slot="header" class="clearfix">
            <span>{{ val.appointingStatus }}</span>
            <el-button v-if="val.appointingStatus === 'appointing'" style="float: right; padding: 3px 0" type="text" @click="updateAppointing(val)">cancel</el-button>
          </div>
          <div class="item">
            <el-tag>squence ID: </el-tag>
            <span>{{ val.objectOfAppointing }}</span>
          </div>
          <div class="item">
            <el-tag type="success">patient ID: </el-tag>
            <span>{{ val.patient }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">hospital ID: </el-tag>
            <span>{{ val.hospital }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">create time: </el-tag>
            <span>{{ val.createTime }}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAppointingList, updateAppointing } from '@/api/appointing'

export default {
  name: 'AppointingPatient',
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
    queryAppointingList({ patient: this.accountId }).then(response => {
      if (response !== null) {
        this.appointingList = response
      }
      this.loading = false
    }).catch(_ => {
      this.loading = false
    })
  },
  methods: {
    updateAppointing(item) {
      this.$confirm('If you need to cancel appointing?', 'tips', {
        confirmButtonText: 'confirm',
        cancelButtonText: 'cancel',
        type: 'success'
      }).then(() => {
        this.loading = true
        updateAppointing({
          patient: item.patient,
          hospital: item.hospital,
          objectOfAppointing: item.objectOfAppointing,
          status: 'cancelled'
        }).then(response => {
          this.loading = false
          if (response !== null) {
            this.$message({
              type: 'success',
              message: 'operation success!'
            })
          } else {
            this.$message({
              type: 'error',
              message: 'operation failed!'
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
          message: 'cancelled'
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

  .d-me-card {
    width: 280px;
    height: 300px;
    margin: 18px;
  }
</style>
