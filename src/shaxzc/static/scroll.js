$(document).ready(function(){
	
	lunboInit();

});

function lunboInit(){
	var yunx;       // 定义动画变量
	var Time = 5000;      // 自动播放间隔时间 毫秒
	var num = 1500;       // 切换图片间隔的时间 毫秒
	var page = 0;         // 定义变量
	var len = $( ".zfb_datu_ul li" ).length;     // 获取图片的数量
	$( ".zfb_datu_ul li" ).css( "opacity", 0 );   // 设置全部的图片透明度为0
	$( ".zfb_datu_ul li:first" ).css( "opacity", 1 ); // 设置默认第一张图片的透明度为1
	function fun(){      // 封装
		$( ".xiaod_div span" ).eq( page ).addClass( "a_active" ).siblings().removeClass( "a_active" );       // 切换小点
		$( ".zfb_datu_ul li" ).eq( page ).animate({ "opacity" : 1 }, num ).siblings().animate( { "opacity" : 0 }, num );     // 切换图片
	
		$( ".zfb_datu_ul li" ).eq( page ).css({ "z-index" :"400" }, num ).siblings().css( { "z-index" :"300" }, num );
	}
	function run(){         // 封装
		if( !$(".zfb_datu_ul li" ).is( ":animated" )){    // 判断li是否在动画
			if( page == len - 1 ){ // 当图片切换到了最后一张的时候
				page = 0;    // 把page设置成0 又重新开始播放动画
				fun();
			}else{     // 继续执行下一张
				page++;
				fun();
			}
		}
	}
	$(".xiaod_div span").click( function(){  // 点击小点
		if( !$( ".zfb_datu_ul li" ).is( ":animated" ) ){   // 判断li是否在动画
			var index = $( ".xiaod_div span" ).index( this );   // 获取点击小点的位置
			page = index;    // 同步于page
			fun();
		}
	});
	
	$( ".ownerBanner" ).hover( function(){    // 鼠标放上去图片的时候清除动画
		clearInterval( yunx );
	}, function(){     // 鼠标移开图片的时候又开始执行动画
		yunx = setInterval( run, Time );
	}).trigger( "mouseleave" );
}

