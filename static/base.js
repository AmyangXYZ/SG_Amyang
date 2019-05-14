String.prototype.format = function(args) {
    var result = this;
    if (arguments.length < 1) {
        return result;
    }

    var data = arguments;       //如果模板参数是数组
    if (arguments.length == 1 && typeof (args) == "object") {
        //如果模板参数是对象
        data = args;
    }
    for (var key in data) {
        var value = data[key];
        if (undefined != value) {
            result = result.replace("{" + key + "}", value);
        }
}
    return result;
}

html = `
    <article class="post">
        <header class="entry-header">
            <div class="entry-title">
                <a href="/posts/{title_url}">{title}</a>
            </div>
            <div class="entry-meta">
                <div class="meta-item">
                    By Amyang
                </div>
                <div class="meta-item">
                    In 
                    <a href="/posts/category/{cat1}">{cat2}</a>
                </div>
                <div class="meta-item-time">
                    {time}
                </div>
            </div>
        </header>
        <div class="entry-content">
            {content}
        </div>
        <div class="entry-footer">
            <a class="button" href="/posts/{readon_url}">Read on&nbsp;&nbsp;<i class="fa fa-angle-double-right"></i></a>
        </div>
    </article>
`
page_home = 1;
page_cat = 1;

$(function() {
    $(window).scroll(function(){
        h = $(window).scrollTop();
        if (h<670) {    
            $("#headerbar").css("top", 0);
            $("#headerbar").css("height","110px");
	    $(".slot-l").css("height","110px");
            $("#headerbar").css("background", "transparent");
        }
        if (h>670 && h<1000) {
            $("#headerbar").css("top", "-99px");
            $("#headerbar").css("background", "transparent");
        }
        if (h>1000) {
            $("#headerbar").css("top", 0);
            $("#headerbar").css("height","70px");
            $(".slot-l").css("height","70px");
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
            if (result.flag == 1) {
                localStorage.setItem("SG_Token", result.data.SG_Token);
                alert("Welcome Home, my Master");
                $("#signin-box").hide();
                $(".overlay").hide();
                $("body").css("overflow", "");
            } else {
                alert(result.msg)
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

    $("#loadmore-home").click(function(){
        page_home+=1;
        $.ajax({
            type: "GET",
            url: "/api/posts/page/"+page_home,
            dataType: "json",
            success:function(result) {
                posts=result.data;
                if (posts.length>0) {
                     for (var i=0; i<posts.length; i++) {
                        h = html.format({"title_url":posts[i].title.replace(/ /g,"-"), 
                        "title":posts[i].title, "cat1":posts[i].cat, "cat2":posts[i].cat, 
                        "time":posts[i].time, "content":posts[i].html.split("<p><i class=\"fa fa-tag fa-emoji\" title=\"tag\"></i></p>")[0],
                        "readon_url":posts[i].title.replace(/ /g,"-")})
                        $(".pagination").before(h);
                     }
                } else {
                    alert("no more")
                }
            },
            error:function(result) {
                alert(result.status);
            }
        });
    })

    $("#loadmore-cat").click(function(){
        page_cat+=1;
        cat = window.location.pathname.split("/")[3]
        $.ajax({
            type: "GET",
            url: "/api/posts/category/{0}/page/{1}".format(cat, page_cat),
            dataType: "json",
            success:function(result) {
                posts=result.data;
                if (posts.length>0) {
                     for (var i=0; i<posts.length; i++) {
                        h = html.format({"title_url":posts[i].title.replace(/ /g,"-"), 
                        "title":posts[i].title, "cat1":posts[i].cat, "cat2":posts[i].cat, 
                        "time":posts[i].time, "content":posts[i].html.split("<p><i class=\"fa fa-tag fa-emoji\" title=\"tag\"></i></p>")[0]})
                        $(".pagination").before(h);
                     }
                } else {
                    alert("no more")
                }
            },
            error:function(result) {
                alert(result.status);
            }
        });
    })
});


