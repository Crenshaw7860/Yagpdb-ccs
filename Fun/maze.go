{{/*Made by Crenshaw#7860
Based off Raffael#7777's Carl-Bot command here: https://carl.gg/t/260353
Command type: prefix/mention
It also shows the solution, and works with .ExecData*/}}
 
{{$seed:=randInt 100000000 999999999}}
{{$embed:=sdict "thumbnail" (sdict "url" (.User.AvatarURL "128")) "title" (title "maze") "description" (joinStr "" "Requested by " .User.String " -") "image" (sdict "url" nil) "color" 123}}
 
{{if .CmdArgs}}
     {{$cross:=(index .CmdArgs 0)}}
     {{$maze:=(joinStr "" "http://maze5.de/cgi-bin/maze?sample=1&type=4&rows=12&columns=12&crossings=" (index .CmdArgs 0) "&seed=" $seed "&algorithm=backtracker&algorithm=0.5&foreground=%23ffffff&background=%2336393f&bordersize=16&cellsize=32&linewidth=2.5&format=png")}}
    {{$embed.image.Set "url" $maze}}
    {{$maze:= (joinStr "" $maze "&solution=true")}}
    {{$embed.Set "description" (print $embed.description " [Solution](" $maze ")")}}
 
{{else if .ExecData}}
   {{$maze:=(joinStr "" "http://maze5.de/cgi-bin/maze?sample=1&type=4&rows=12&columns=12&crossings=" .ExecData "&seed=" $seed "&algorithm=backtracker&algorithm=0.5&foreground=%23ffffff&background=%2336393f&bordersize=16&cellsize=32&linewidth=2.5&format=png")}}
    {{$embed.image.Set "url" $maze}}
    {{$maze:= (joinStr "" $maze "&solution=true")}}
    {{$embed.Set "description" (print $embed.description " [Solution](" $maze ")")}}
{{else}}
   {{$maze:=(joinStr "" "http://maze5.de/cgi-bin/maze?sample=1&type=4&rows=12&columns=12&crossings=0&seed=" $seed "&algorithm=backtracker&algorithm=0.5&foreground=%23ffffff&background=%2336393f&bordersize=16&cellsize=32&linewidth=2.5&format=png")}}
    {{$embed.image.Set "url" $maze}}
    {{$maze:= (joinStr "" $maze "&solution=true")}}
    {{$embed.Set "description" (print $embed.description " [Solution](" $maze ")")}}
{{end}}
{{sendMessage nil (cembed $embed)}}