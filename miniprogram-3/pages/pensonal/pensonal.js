// pages/pensonal/pensonal.js
Page({

  data: {
    showModal: false // 控制模态框是否显示
  },
  // 函数用于显示模态框
  showModal: function() {
    this.setData({ showModal: true });
  },
  // 函数用于隐藏模态框
  hideModal: function() {
    this.setData({ showModal: false });
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var that = this; // 保留当前上下文的引用
    wx.getSystemInfo({
      success: function (res) {
        that.setData({
          statusBarHeight: res.statusBarHeight // 更新状态栏高度
        });
      }
    });
  },
  showActionSheet: function () {
    const that = this;
    wx.showActionSheet({
      itemList: ['登录', '注册'], // 操作菜单列表
      success: function (res) {
        if (res.tapIndex === 0) {
          // 用户选择了登录
          // 转到登录界面，或者直接展开登录表单
          wx.navigateTo({
            url: '/pages/login/login' // 页面路径要根据自己的项目设置
          });
        } else if (res.tapIndex === 1) {
          // 用户选择了注册
          // 转到注册界面，或者直接展开注册表单
          wx.navigateTo({
            url: '/pages/register/register' // 页面路径要根据自己的项目设置
          });
        }
      },
      fail: function (res) {
        console.log(res.errMsg);
      }
    });
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  }
})