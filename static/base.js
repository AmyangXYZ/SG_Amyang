$(document).ready(function(){
    $(".sidebar").hide();
    $(".sidebar-overlay").hide();
    $("#sidebar-button").click(function(){
        $("body").css("overflow", "hidden");
        $(".sidebar").show(300);
        $(".sidebar-overlay").show();
    })
    $(".sidebar-overlay").click(function(){
        $("body").css("overflow", "");
        $(".sidebar").hide(300);
        $(".sidebar-overlay").hide();
    })
});