# flat file og krypteringsopgave

## Implementation

**Krypterings- og hashing-algoritme valg**
*Kryptering* - da dataen der skal krypteres skal "gemmes" er standarden at bruge symmetrisk kryptering da det er hurtigt, der stod valget på AES, med 256-bit nøgle, som automatisk bliver valgt i Go når nøglen er præcis 32 byte lang.

*Hashing* - Valget stod på bcrypt da den kan modstå "brute-force" angreb, samt at den var let tilgængelig i Go.

**Hvornår og hvorfor kryptere**
*Hvornår* - Data skal krypteres når det skal gemmes på en harddisk, database eller lignende og når det skal sendes over netværket, som tit bliver gjort over HTTPS.

*Hvorfor* - Når vi følger CIA modellen, så er en af punkterne fortrolighed, her vil vi gerne sikre mod at uvedkommende ikke kan se dataen, hvis de for adgang til dataen. Der er også GDPR, som via lovkrav kræver kryptering af personfølsomme oplysninger.

**Hvornår og hvorfor dekryptere**
*Hvornår* - data skal kun dekrypteres når en autoriseret bruger eller system skal bruge dataen. 

*Hvorfor* - For at se den reelle data og ikke den krypteret data skal vi dekryptere.

**Hvornår og hvorfor fjernes dekrypteret data fra hukommelsen**
*Hvornår* - Så snart at dataen ikke skal bruges længere, skal den slettes eller overskrives i RAM

*Hvorfor* - Data kan ligge i RAM i et stykke tid inden en GC fjerner det, så derfor er det vigtigt at enten overskrive eller slette det, så snart det ikke skal bruges længere. Dette beskytter mod Memory dumping.


## Resultat

