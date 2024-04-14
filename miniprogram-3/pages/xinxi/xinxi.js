// pages/xinxi/xinxi.js
Page({

  /**
   * 页面的初始数据
   */
  data: {

  },
// 提交表单的事件处理函数
  submitForm: function(e) {
    let formData = e.detail.value; 
    console.log(formData);
    wx.request({
      url: 'https://127.0.0.1/add_competition', 
      method: 'POST',
      data: formData,
      success: function(res) {
        if (res.statusCode == 200) {
          wx.showToast({
            title: '提交成功',
            icon: 'success'
          });
          wx.navigateTo({
            url: '/pages/home/home'
          });
        } else {
          wx.showToast({
            title: '提交失败',
            icon: 'none'
          });
        }
      }
    });
  },
  // 处理取消操作的事件处理函数
  cancelForm: function() {
    wx.navigateBack({
      delta: 1 
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    var that = this;
    wx.getSystemInfo({
      success: function (res) {
        that.setData({
          statusBarHeight: res.statusBarHeight 
        });
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