import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [{
  path: '/login',
  component: () => import('@/views/login/index'),
  hidden: true
},

{
  path: '/404',
  component: () => import('@/views/404'),
  hidden: true
},

{
  path: '/',
  component: Layout,
  redirect: '/realSequence',
  children: [{
    path: 'realSequence',
    name: 'RealSequence',
    component: () => import('@/views/realSequence/list/index'),
    meta: {
      title: 'DNA Information',
      icon: 'realSequence'
    }
  }]
}
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: '/authorizing',
    component: Layout,
    redirect: '/authorizing/all',
    name: 'Authorizing',
    alwaysShow: true,
    meta: {
      title: 'Authorizing',
      icon: 'authorizing'
    },
    children: [{
      path: 'all',
      name: 'AuthorizingAll',
      component: () => import('@/views/authorizing/all/index'),
      meta: {
        title: 'Authorization Information',
        icon: 'authorizingAll'
      }
    },
    {
      path: 'me',
      name: 'AuthorizingMe',
      component: () => import('@/views/authorizing/me/index'),
      meta: {
        roles: ['editor'],
        title: 'Authorizing by Me',
        icon: 'authorizingMe'
      }
    }, {
      path: 'authorize',
      name: 'AuthorizingBuy',
      component: () => import('@/views/authorizing/authorize/index'),
      meta: {
        roles: ['editor'],
        title: 'Authorized',
        icon: 'authorizingBuy'
      }
    }
    ]
  },
  {
    path: '/appointing',
    component: Layout,
    redirect: '/appointing/all',
    name: 'Appointing',
    alwaysShow: true,
    meta: {
      title: 'Appointing',
      icon: 'appointing'
    },
    children: [{
      path: 'all',
      name: 'AppointingAll',
      component: () => import('@/views/appointing/all/index'),
      meta: {
        title: 'Appointing information',
        icon: 'appointingAll'
      }
    },
    {
      path: 'patient',
      name: 'AppointingPatient',
      component: () => import('@/views/appointing/patient/index'),
      meta: {
        roles: ['editor'],
        title: 'Appointing by me',
        icon: 'appointingPatient'
      }
    }, {
      path: 'Hospital',
      name: 'AppointingHospital',
      component: () => import('@/views/appointing/hospital/index'),
      meta: {
        roles: ['editor'],
        title: 'Appointed',
        icon: 'appointingHospital'
      }
    }
    ]
  },
  {
    path: '/addRealSequence',
    component: Layout,
    meta: {
      roles: ['admin']
    },
    children: [{
      path: '/addRealSequence',
      name: 'AddRealSequence',
      component: () => import('@/views/realSequence/add/index'),
      meta: {
        title: 'Add DNA Information',
        icon: 'addRealSequence'
      }
    }]
  },

  // 404 page must be placed at the end !!!
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

const createRouter = () => new Router({
  base: '/web',
  // mode: 'history', // require service support
  scrollBehavior: () => ({
    y: 0
  }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
