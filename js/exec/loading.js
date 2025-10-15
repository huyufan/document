// jq版本的上拉刷新下拉加载组件
(function ($) {
  // 防抖函数
  function debounce(func, wait, immediate) {
    var timeout;

    return function () {
      var context = this,
        args = arguments;
      var later = function () {
        timeout = null;
        if (!immediate) func.apply(context, args);
      };
      var callNow = immediate && !timeout;
      clearTimeout(timeout);
      timeout = setTimeout(later, wait);
      if (callNow) func.apply(context, args);
    };
  }

  $.fn.pullToLoading = function (onPullDown, onPullUp) {
    var startY = 0,
      moveY = 0;
    var debouncedPullDown = debounce(onPullDown, 500, true);
    var debouncedPullUp = debounce(onPullUp, 500, true);

    this.on({
      touchstart: function (e) {
        startY = e.touches[0].pageY;
      },
      touchmove: function (e) {
        moveY = e.touches[0].pageY;
        if (this.scrollTop === 0 && moveY > startY) {
          e.preventDefault();
          if (debouncedPullDown) debouncedPullDown();
        } else if (this.scrollTop + this.clientHeight >= this.scrollHeight) {
          if (debouncedPullUp) debouncedPullUp();
        }
      },
    });
    return this;
  };
})(jQuery);

// 使用方式
$(".scroll-box").pullToLoading(
  function () {
    // renderList();
    // 下拉刷新操作
    console.log("下拉成功----");
  },
  function () {
    // 上拉加载操作
    console.log("上拉成功----");
    // renderList();
  }
);
