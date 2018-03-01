$(function() {
        var editor = editormd("editormd", {
        path : "/static/editor-md/lib/",
        width: 1050,
        height: 640,
        emoji: true,
        saveHTMLToTextarea : true,
        codeFold : true,
    });

    $('#NewPostForm').submit(function (event) {
        event.preventDefault();
        var form = $(this);
        data = {"title":$("#title").val(), "cat":$("#cat").val(), "html":editor.getHTML(), "md":editor.getMarkdown()}
        console.log(data);
        title = $("#title").val().replace(" ","-");
        $.ajax({
            async: false,
            type: "POST",
            url: form.attr('action'),
            data: data,
            dataType: "JSON",
            headers: {
                "Authorization": "SG_Token "+localStorage.getItem("SG_Token")
            }
        }).done(function(result) {
            if(result.status=="success"){
                alert("Created !");
                $(location).attr('href', '/posts/'+title);
            } else {
                alert(result.message);
            }
        }).fail(function(){
            alert("where is ur token?")
        });
    });

    $('#UpdatePostForm').submit(function (event) {
        event.preventDefault();
        var form = $(this);
        data = {"new-title":$("#title").val(), "cat":$("#cat").val(), "html":editor.getHTML(), "md":editor.getMarkdown()}
        title = $("#title").val();
        $.ajax({
            async: false,
            type: "PUT",
            url: form.attr('action')+"/"+title,
            data: data,
            dataType: "JSON",
            headers: {
                "Authorization": "SG_Token "+localStorage.getItem("SG_Token")
            }
        }).done(function(result) {
            if(result.status=="success"){
                alert("Updated!");
                $(location).attr('href', '/posts/'+title.replace(" ","-"));
            } else {
                alert(result.message);
            }
        }).fail(function(){
            alert("where is ur token?")
        });
    });

});