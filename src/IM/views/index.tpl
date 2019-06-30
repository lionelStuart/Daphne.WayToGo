{{template "base/base.html" .}}
{{define "head"}}
<title>{{i18n .Lang "welcome"}} - {{i18n .Lang "app_intro"}}</title>
{{end}}
{{define "body"}}

<form action="/join" method="post" class="form-horizontal">
    <div class="form-group">
        <label class="col-md-3 control-label">{{i18n .Lang "username"}}: </label>
        <div class="col-md-5">
              <input type="text" class="form-control" name="uname" required>
        </div>
    </div>
    <div class="form-group">
        <label class="col-md-3 control-label">{{i18n .Lang "technology"}}: </label>
        <div class="col-md-5">
            <select class="form-control" name="tech">
                <option value="longpolling">{{i18n .Lang "longpolling"}}</option>
                <option value="websocket">{{i18n .Lang "websocket"}}</option>
            </select>
        </div>
    </div>

        <div class="form-group">
            <label class="col-md-3 control-label">{{i18n .Lang "roomid"}}: </label>
            <div class="col-md-5">
                <select class="form-control" name="roomId">
                    {{range .goods}}
                     <option value="{{ .}}">{{ .}}</option>
                    {{ end  }}
                     <option value="0">{{i18n .Lang "new"}}</option>
                </select>
            </div>
        </div>

    <div class="form-group">
        <div class="col-sm-offset-3 col-sm-10">
            <button type="submit" class="btn btn-info">{{i18n .Lang "enter_chat"}}</button>
        </div>
    </div>
</form>
{{end}}