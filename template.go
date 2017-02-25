package main

const irssiConfigTemplate = `# Generated by yarssi
servers = ({{range .Servers}}
  {
    address = "{{.Address}}";
    chatnet = "{{.Name}}";
    port = "{{.Port}}";
    autoconnect = "{{if .Autoconnect}}yes{{else}}no{{end}}";
    {{if .Password}}password = "{{.Password}}";{{end}}
    use_ssl = "{{if .Ssl}}yes{{else}}no{{end}}";
  },{{end}}
);

chatnets = { {{range .Servers}}
  {{.Name}} = {
    type = "IRC";
    max_kicks = "4";
    max_msgs = "5";
    max_whois = "4";
    max_query_chans = "5";
    {{if .Autosendcmd }}autosendcmd = "{{.Autosendcmd}}";{{end}}
  };{{end}}
};


channels = ( {{range .Windows}}{{if .IsChannel }}
  { name = "{{.Name}}"; chatnet = "{{.Server}}"; autojoin = "Yes"; }, {{end}}{{end}}
);

aliases = {
  J = "join";
  WJOIN = "join -window";
  WQUERY = "query -window";
  LEAVE = "part";
  BYE = "quit";
  EXIT = "quit";
  SIGNOFF = "quit";
  DESCRIBE = "action";
  DATE = "time";
  HOST = "userhost";
  LAST = "lastlog";
  SAY = "msg *";
  WI = "whois";
  WII = "whois $0 $0";
  WW = "whowas";
  W = "who";
  N = "names";
  M = "msg";
  T = "topic";
  C = "clear";
  CL = "clear";
  K = "kick";
  KB = "kickban";
  KN = "knockout";
  BANS = "ban";
  B = "ban";
  MUB = "unban *";
  UB = "unban";
  IG = "ignore";
  UNIG = "unignore";
  SB = "scrollback";
  UMODE = "mode $N";
  WC = "window close";
  WN = "window new hide";
  SV = "say Irssi $J ($V) - http://irssi.org/";
  GOTO = "sb goto";
  CHAT = "dcc chat";
  RUN = "SCRIPT LOAD";
  CALC = "exec - if command -v bc >/dev/null 2>&1\\; then printf '%s=' '$*'\\; echo '$*' | bc -l\\; else echo bc was not found\\; fi";
  SBAR = "STATUSBAR";
  INVITELIST = "mode $C +I";
  Q = "QUERY";
  "MANUAL-WINDOWS" = "set use_status_window off;set autocreate_windows off;set autocreate_query_level none;set autoclose_windows off;set reuse_unused_windows on;save";
  EXEMPTLIST = "mode $C +e";
  ATAG = "WINDOW SERVER";
  UNSET = "set -clear";
  RESET = "set -default";
  wg = "/window goto";
  hideadd = "eval set activity_hide_targets $activity_hide_targets $-";
};
statusbar = {
  # formats:
  # when using {templates}, the template is shown only if it's argument isn't
  # empty unless no argument is given. for example {sb} is printed always,
  # but {sb $T} is printed only if $T isn't empty.

  items = {
    # start/end text in statusbars
    barstart = "{sbstart}";
    barend = "{sbend}";

    topicbarstart = "{topicsbstart}";
    topicbarend = "{topicsbend}";

    # treated "normally", you could change the time/user name to whatever
    time = "{sb $Z}";
    user = "{sb {sbnickmode $cumode}$N{sbmode $usermode}{sbaway $A}}";

    # treated specially .. window is printed with non-empty windows,
    # window_empty is printed with empty windows
    window = "{sb $winref:$tag/$itemname{sbmode $M}}";
    window_empty = "{sb $winref{sbservertag $tag}}";
    prompt = "{prompt $[.15]itemname}";
    prompt_empty = "{prompt $winname}";
    topic = " $topic";
    topic_empty = " Irssi v$J - http://www.irssi.org";

    # all of these treated specially, they're only displayed when needed
    lag = "{sb Lag: $0-}";
    act = "{sb Act: $0-}";
    more = "-- more --";
  };

  # there's two type of statusbars. root statusbars are either at the top
  # of the screen or at the bottom of the screen. window statusbars are at
  # the top/bottom of each split window in screen.
  default = {
    # the "default statusbar" to be displayed at the bottom of the window.
    # contains all the normal items.
    window = {
      disabled = "no";

      # window, root
      type = "window";
      # top, bottom
      placement = "bottom";
      # number
      position = "1";
      # active, inactive, always
      visible = "active";

      # list of items in statusbar in the display order
      items = {
        barstart = { priority = "100"; };
        time = { };
        user = { };
        window = { };
        window_empty = { };
        lag = { priority = "-1"; };
        more = { priority = "-1"; alignment = "right"; };
        barend = { priority = "100"; alignment = "right"; };
      };
    };

    # statusbar to use in inactive split windows
    window_inact = {
      type = "window";
      placement = "bottom";
      position = "1";
      visible = "inactive";
      items = {
        barstart = { priority = "100"; };
        window = { };
        window_empty = { };
        more = { priority = "-1"; alignment = "right"; };
        barend = { priority = "100"; alignment = "right"; };
      };
    };

    # we treat input line as yet another statusbar :) It's possible to
    # add other items before or after the input line item.
    prompt = {
      type = "root";
      placement = "bottom";
      # we want to be at the bottom always
      position = "100";
      visible = "always";
      items = {
        prompt = { priority = "-1"; };
        prompt_empty = { priority = "-1"; };
        # treated specially, this is the real input line.
        input = { priority = "10"; };
      };
    };

    # topicbar
    topic = {
      type = "root";
      placement = "top";
      position = "1";
      visible = "always";
      items = {
        topicbarstart = { priority = "100"; };
        topic = { };
        topic_empty = { };
        topicbarend = { priority = "100"; alignment = "right"; };
      };
    };
    awl_0 = {
      items = {
        barstart = { priority = "100"; };
        awl_0 = { };
        barend = { priority = "100"; alignment = "right"; };
      };
    };
    awl_1 = {
      items = {
        barstart = { priority = "100"; };
        awl_1 = { };
        barend = { priority = "100"; alignment = "right"; };
      };
    };
    awl_2 = {
      items = {
        barstart = { priority = "100"; };
        awl_2 = { };
        barend = { priority = "100"; alignment = "right"; };
      };
    };
    awl_3 = {
      items = {
        barstart = { priority = "100"; };
        awl_3 = { };
        barend = { priority = "100"; alignment = "right"; };
      };
    };
    awl_4 = {
      items = {
        barstart = { priority = "100"; };
        awl_4 = { };
        barend = { priority = "100"; alignment = "right"; };
      };
    };
    awl_5 = {
      items = {
        barstart = { priority = "100"; };
        awl_5 = { };
        barend = { priority = "100"; alignment = "right"; };
      };
    };
    awl_6 = {
      items = {
        barstart = { priority = "100"; };
        awl_6 = { };
        barend = { priority = "100"; alignment = "right"; };
      };
    };
  };
};
settings = {
  core = {
    real_name = "{{.Realname}}";
    user_name = "{{.Username}}";
    nick = "{{.Nickname}}";
    timestamp_format = "%H:%M:%S";
  };
  "fe-text" = { actlist_sort = "refnum"; };
  "fe-common/core" = {
    theme = "irssilteplait.theme";
    activity_hide_level = "QUITS JOINS PARTS MODES NICKS TOPIC";
    term_charset = "UTF-8";
    autolog = "yes";
    activity_hide_targets = "{{range .Windows}}{{if not .Show}} {{.Name}}{{end}}{{end}}";
    autolog_ignore_targets = "";
  };
  "perl/core/scripts" = {
    awl_shared_sbar = "OFF";
    awl_viewer = "no";
    awl_hide_data = "1";
    awl_height_adjust = "2";
    screen_away_message = "afk";
    autoaway_timeout = "600";
    spellcheck_languages = "#sinf/fr_CH";
    spellcheck_enabled = "no";
  };
};
hilights = ({{range .Hilights}}
  { text = "{{.}}"; nick = "yes"; word = "yes"; },{{end}}
);
logs = { };
ignores = (
  { mask = "suuuuyi"; level = "ALL"; },
  { mask = "MeetBot"; level = "ALL"; }
);
windows = {
  1 = { immortal = "yes"; name = "(status)"; level = "ALL"; };
  {{range $i, $v := .Windows}}
  {{ $v.Inc $i  }} = {
    items = (
      {
        type = "{{if $v.IsChannel}}CHANNEL{{else}}QUERY{{end}}";
        chat_type = "IRC";
        name = "{{$v.Name}}";
        tag = "{{$v.Server}}";
      }
    );
  };
  {{ end }}
};
`
