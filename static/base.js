$(document).ready(function(){
    $(window).scroll(function() {
        $(".slot-l").css("opacity",1-$(window).scrollTop()/500);
    });
});