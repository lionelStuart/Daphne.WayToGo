var socket;

$(document).ready(function () {
    // Create a socket
    socket = new WebSocket('ws://' + window.location.host + '/ws/join?uname=' + $('#uname').text()+'&roomid='+$('#roomid').text());
    console.log('ws://' + window.location.host + '/ws/join?uname=' + $('#uname').text()+'&roomId='+$('#roomid').text())
    // Message received on the socket
    socket.onmessage = function (event) {
        var data = JSON.parse(event.data);
        var li = document.createElement('li');

        console.log(data);

        switch (data.Type) {
        case 0: // JOIN
            if (data.UserInfo.Uname == $('#uname').text()) {
                li.innerText = 'You joined the chat room.';
            } else {
                li.innerText = data.UserInfo.Uname + ' joined the chat room.';
            }
            if(data.UserInfo.RoomId !=$('#roomid').text()){
                console.log("roomid is different ",data.UserInfo.RoomId,$('#roomid').text())
                document.getElementById("roomid").innerHTML=data.UserInfo.RoomId
            }
            break;
        case 1: // LEAVE
            li.innerText = data.User + ' left the chat room.';
            break;
        case 2: // MESSAGE
            var username = document.createElement('strong');
            var content = document.createElement('span');

            username.innerText = data.UserInfo.Uname;
            content.innerText = data.Content;

            li.appendChild(username);
            li.appendChild(document.createTextNode(': '));
            li.appendChild(content);

            break;
        }

        $('#chatbox li').first().before(li);
    };

    // Send messages.
    var postConecnt = function () {
        var uname = $('#uname').text();
        var content = $('#sendbox').val();
        socket.send(content);
        $('#sendbox').val('');
    }

    $('#sendbtn').click(function () {
        postConecnt();
    });
});
