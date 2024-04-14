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

  inputChange: function(e) {
    this.setData({
      buttonName: e.detail.value // 更新按钮名称为用户输入的值
    });
  },

  // 显示输入框
  addButton: function() {
    this.setData({
      showInput: true // 点击加号按钮后显示输入框
    });
  },

  // 实际添加按钮的逻辑
  confirmButton: function() {
    const that = this;
      wx.request({
      url: 'https://127.0.0.1/addCompetotion', 
      method: 'POST',
      data: {
        buttonName: this.data.buttonName // 发送输入的按钮名称
      },
      header: {
        'content-type': 'application/json' // 设置请求的 header
      },
      success(res) {
        // 请求成功的处理
        if (res.statusCode == 200) {
          // 这里编写请求成功后的逻辑，例如提示添加成功，清空输入框等
          wx.showToast({
            title: '添加成功',
            icon: 'success',
            duration: 2000
          });

          // 清空输入框
          that.setData({
            buttonName: '',
            showInput: false 
          });
        } else {
          // 处理请求失败的情况
          wx.showToast({
            title: '添加失败',
            icon: 'none',
            duration: 2000
          });
        }
      },
      fail() {
        // 请求发送失败，可能是网络错误等原因
        wx.showToast({
          title: '请求失败',
          icon: 'none',
          duration: 2000
        });
      }
    });
  },

  cancelAddButton: function() {
    this.setData({
      showInput: false, // 隐藏输入框
      buttonName: '' // 可选：重置输入框的内容
    });
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