# Træning af model

Træningsdataen i form af .jpg billeder skal befinde sig i `data/<genstand>/`.

Hvis man har taget en video, så kan man få billeder fra optagelsen med

```bash
$ ffmpeg -i input.mp4 -vf fps=2 %04d.png
```

hvor `fps=2` betyder, at der skal tages to billeder fra hvert sekund.

Navnene burde være ligegyldige.

# Konvert til JS
Hent tensorflowjs med conda, pip eller hvad man nu har lyst til. Kør
```bash
tensorflowjs_converter \
    --input_format=tf_saved_model \
    --output_node_names='MobilenetV1/Predictions/Reshape_1' \
    --saved_model_tags=serve \
    /fitted_model_smaller \
    /web_model
```

Det lader til, at output_node_names ikke er særlig vigtig. Beskrivelsen
af indstillingen er
> The names of the output nodes, separated by commas.

Det siger ikke særlig meget. `MobilenetV1` er en rigtig model, hvor
der også findes v2 og v3. Hvis det har en betydning, så vil jeg jo
hellere bruge en nyere udgave.
