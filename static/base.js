$(function() {

    $(window).scroll(function(){
        h = $(window).scrollTop();
        if (h<670) {    
            $("#headerbar").css("top", 0);
            $("#headerbar").css("height","110px");
            $("#headerbar").css("background", "transparent");
        }
        if (h>670 && h<1000) {
            $("#headerbar").css("top", "-99px");
            $("#headerbar").css("background", "transparent");
        }
        if (h>1000) {
            $("#headerbar").css("top", 0);
            $("#headerbar").css("height","70px");
            $("#headerbar").css("background", "black");
        }
    })

    // show line number for hljs
    $('code.hljs').each(function(i, block) {
        hljs.lineNumbersBlock(block);
    });
    
    // sidebar
    $("#sidebar-open").click(function(){
        $(".sidebar").css("right",415);
        $("body").css("overflow", "hidden");
        $(".sidebar").show(300);
        $(".overlay").show();
    })
    $(".overlay").click(function(){
        $("#signin-box").hide();
        $("#upload-box").hide();
        $("body").css("overflow", "");
        $(".sidebar").css("right",0);
        $(".overlay").hide();
    })
    $("#sidebar-close").click(function(){
        $("body").css("overflow", "");
        $(".sidebar").css("right",0);
        $(".overlay").hide();
    })

    // Sign In Box
    $("#sub-menu-content-item-signin").click(function(){
        $("#signin-box").show();
        $(".overlay").show();
        $("body").css("overflow", "hidden");
    })

    $('#SignInForm').submit(function (event) {
        event.preventDefault();
        var form = $(this);
        $.ajax({
            async: false,
            type: form.attr('method'),
            url: form.attr('action'),
            data: form.serialize(),
            dataType: "JSON",
        }).done(function (result) {
            if (result.status == "success") {
                localStorage.setItem("SG_Token", result.data.SG_Token);
                alert("Welcome Home, my Master");
                $("#signin-box").hide();
                $(".overlay").hide();
                $("body").css("overflow", "");
            } else {
                alert(result.message)
            }
        }).fail(function(result){
            alert("some thing error")
        });
    });

    // Upload Box
    $("#sub-menu-content-item-upload").click(function(){
        $("#upload-box").show();
        $(".overlay").show();
        $("body").css("overflow", "hidden");
    })

    $("#UploadForm").submit(function (event) {
        event.preventDefault();
        var form = $(this);
        var formData = new FormData(this);
        $.ajax({
            async: false,
            type: "POST",
            url: form.attr('action'),
            data: formData,
            mimeType: "multipart/form-data",
            contentType: false,
            cache: false,
            processData: false,
            dataType: "JSON",
            headers: {
                "Authorization": "SG_Token "+localStorage.getItem("SG_Token")
            }
        }).done(function(result) {
            $("#filePath").html(result.data);
        }).fail(function() {
            alert("where is ur token?");
        })
    });

    $("#loadmore").click(function(){
        $.ajax({
            type: "POST",
            url: "/api/posts",
            headers: {
                "Authorization": "SG_Token "+localStorage.getItem("SG_Token")
            },
            dataType: "json",
            success:function(result) {
                post=result.data;
                alert(result.status);
            },
            error:function(result) {
                alert(result.status);
            }
        });
    })
});