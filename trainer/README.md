# Træning af model
Træningsdataen i form af .jpg billeder skal befinde sig i `data/<genstand>/`.

Hvis man har taget en video, så kan man få billeder fra optagelsen med 
```bash
$ ffmpeg -i input.mp4 -vf fps=2 %04d.png
```
hvor `fps=2` betyder, at der skal tages to billeder fra hvert sekund.

Navnene burde være ligegyldige.