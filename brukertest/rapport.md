# Rapport fra Brukertest av "DJ set database app"

## Introduksjon
I denne rapporten analyseres resultatene fra en nylig brukertest av "DJ set database app", en applikasjon designet for å tilby en omfattende database over DJ-sett. Appen er primært rettet mot unge voksne i aldersgruppen 17 til 25 år. Målet med brukertesten var å identifisere potensielle områder for forbedring i brukergrensesnittet og funksjonaliteten, spesielt i hvordan appen håndterer brukerinput gjennom CLI (Command Line Interface).

## Metode
Testen involverte deltagere rekruttert fra en klasse, som alle tilhører målgruppen. Testen fokuserte på å observere brukernes interaksjon med appens hovedfunksjoner og å samle tilbakemeldinger om deres brukeropplevelse. Observasjonene og tilbakemeldingene ble nøye dokumentert for å gi et grunnlag for videre analyse.

## Resultater
Under brukertesten ble det avdekket flere kritiske feil og områder for forbedring, som detaljert nedenfor:

1. **Håndtering av tilfeldig input**
   - **Problem:** Applikasjonen krasjet ved mottak av uventet eller tilfeldig input fra brukerne.
   - **Brukerfeedback:** Dette førte til frustrasjon og forvirring blant testdeltakerne, som forventet en mer robust feilhåndtering.

2. **Brukergrensesnittets klarhet**
   - **Problem:** Listen som presenterer informasjon var ikke alltid i logisk eller intuitiv rekkefølge, noe som gjorde det vanskelig for brukerne å navigere effektivt.
   - **Brukerfeedback:** Flere brukere foreslo en mer oversiktlig og strukturert presentasjon av data, samt tydeligere instruksjoner om hvordan appen skal brukes.

3. **Språk og kommunikasjon**
   - **Problem:** Språket som ble brukt i appens grensesnitt var til tider for komplisert, noe som gjorde det utfordrende for brukere å forstå funksjonalitet og instruksjoner.
   - **Brukerfeedback:** Det ble foreslått å bruke et enklere og mer tilgjengelig språk for å forbedre forståelsen og brukervennligheten.

4. **Informasjon ved oppstart**
   - **Problem:** Ved første gangs bruk av appen var det mangel på informasjon om appens formål og grunnleggende bruk.
   - **Brukerfeedback:** Brukere ønsket en innledende beskrivelse eller guide ved oppstart som forklarer hva appen handler om og hvordan den brukes.

## Diskusjon
De identifiserte problemene fra brukertesten peker tydelig på flere områder hvor applikasjonen kan forbedres for å øke brukervennligheten og stabiliteten:

1. **Robusthet mot tilfeldig input**
   - For å forhindre krasj ved uventet input, bør det implementeres mer omfattende feilhåndtering og validering av input. Dette vil ikke bare forbedre programmets stabilitet, men også brukeropplevelsen ved å tilby tilbakemeldinger om feil bruk av applikasjonen.

2. **Forbedring av brukergrensesnittet**
   - En klarere og mer intuitiv strukturering av brukergrensesnittet kan oppnås gjennom bruk av bedre organisering av informasjon og lister. Dette inkluderer å prioritere informasjon basert på hva brukerne oftest søker etter eller trenger tilgang til.

3. **Klarere språkbruk**
   - Ved å forenkle språket og gjøre instruksjonene mer tilgjengelige, kan appen bli mer tilgjengelig for målgruppen. Dette kan involvere bruk av mindre teknisk jargon og mer hverdagsspråk som er lett forståelig for nye brukere.

4. **Innføring av startguide**
   - En kort introduksjonssekvens eller guide ved første gangs bruk kan dramatisk forbedre brukernes første inntrykk og forståelse av appen. Dette vil hjelpe nye brukere å raskt få en oversikt over appens funksjoner og hvordan de kan navigere seg rundt.

## Anbefalinger
Basert på analysen av brukertesten og de identifiserte problemområdene, foreslås følgende forbedringer:

1. **Implementere avansert inputvalidering**
   - Utvikle og integrere robuste valideringsrutiner som kan håndtere og gi tilbakemeldinger på ulike typer input. Dette vil forbedre programmets stabilitet og redusere risikoen for at det krasjer under bruk.

2. **Reorganisere brukergrensesnittet**
   - Gjennomføre en redesign av brukergrensesnittet for å gjøre det mer intuitivt og brukervennlig. Dette kan inkludere forbedringer som bedre visuell organisering, enklere navigasjon og mer relevante datafremstillinger basert på brukerens behov.

3. **Forenkle språket i applikasjonen**
   - Gjennomgå og revidere tekster innen applikasjonen for å sikre at de er klare og enkle å forstå. Dette vil bidra til å gjøre appen mer tilgjengelig for alle brukere, uavhengig av deres tekniske forståelse.

4. **Utvikle en velkomstguide**
   - Lage en interaktiv guide eller en serie med introduksjonsvideoer som nye brukere kan følge når de første gang starter appen. Guiden bør dekke grunnleggende funksjoner og vanlige brukerspørsmål for å raskt orientere brukerne.

## Konklusjon
Brukertesten av "DJ set database app" har avdekket viktige områder hvor applikasjonen kan forbedres for å møte brukernes forventninger og behov bedre. Ved å implementere de foreslåtte forbedringene vil applikasjonen ikke bare bli mer stabil og brukervennlig, men også mer tiltalende for målgruppen. Videre testing og brukertilbakemeldinger vil være avgjørende for å sikre at applikasjonen fortsetter å utvikle seg i riktig retning.
