{{ with .Sender }}EHLO {{ .Domain }}
MAIL FROM:{{ .MailFrom }}
DATA
From: {{ .Header.FriendlyFrom }}
Subject: {{ .Header.Subject }}
{{ end }}
Hello Lica.

Did you know lorem20?
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque ac eros turpis. 

Your's trully,
Dirdel.
.