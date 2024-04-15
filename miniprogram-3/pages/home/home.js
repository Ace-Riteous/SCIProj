// pages/home/home.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    buttons: [],     
    buttonName: '', 
    showInput: false  
  },

  onShow: function() {
    this.fetchButtons();
  },

  fetchButtons: function() {
    wx.request({
      url: 'https://127.0.0.1/see_competitions', 
      method: 'GET',
      success: (res) => {
        // 假设服务器返回的数据格式是 { data: [{name: '比赛1'}, {name: '比赛2'}] }
        if (res.statusCode === 200) {
          this.setData({ buttons: res.data.data });
        }
      },
      fail: (err) => {
        console.error('请求失败:', err);
      }
    });
  },

  addButton: function() {
    wx.navigateTo({
      url: '/pages/xinxi/xinxi' 
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
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