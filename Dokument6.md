**Obecné zadání projektu**

Serverový chatbot odpovídající na tři dotazy. Uživatel má k dispozici grafické rozhraní přes které odesílá dotazy na server a následné odpovědi dostává zpět do stejného grafického rozhraní. Dotazy umožňují zjistit aktuální kurz eura, aktuální čas na běžícím serveru a jméno chatbota.

**Zadávání vstupu**

Uživatel zadává vstup do jasně určeného a viditelného textového pole. Odeslání vstupu probíhá po stisknutí tlačítka k tomu určenému. Tlačítko je označené textem ,,Send” a po jeho stisknutí je uživatelem zadaný text odeslán pro zpracování na server. Uživatelský vstup musí být zadáván v anglickém jazyce pro zaručení správnosti funkce chatbota.

**Získávání výstupu**

Výstup bude zobrazen na obrazovce v textové podobě. Je zobrazen po zpracování vstupu, jeho následném vyhodnocení a získáním příslušného výsledku. Výstup je v anglickém jazyce a časové formáty podléhají formátu HH:MM:SS, popřípadě DD:MM:YYYY.

**UseCase diagram**

![](img.png)









**Funkce UseCase diagramu**

- **Odeslání dotazu**
  - Odeslání libovolného nespecifikovaného textu na chatbota
  - V případě odeslání vstupního textu obsahujícího klíčová slova pro více příkazů, budou vráceny všechny příkazy v pořadí, v jakém se v textu nacházejí.
  - V případě nerozpoznání či špatného zadání příkazu bude vráceno “Unknown command or poorly specified text”
  - **Získat aktuální kurz eura**
    - Vrátí aktuální kurz eura vůči české koruně dle API …
    - Je zavolán za předpokladu, že vstupní text obsahuje klíčová slova “rate”, “exchange” či “euro”
    - Vrací výstup ve formátu “Actuall price of one € is *aktuální kurz* CZK”
  - **Získat aktuální čas na serveru**
    - Vrátí aktuální čas na serveru na kterém běží chatbot
    - Je zavolán za předpokladu, že vstupní text obsahuje klíčová slova “time” nebo “date”.
    - Vrací výstup ve formátu “Actuall server time is *aktuální čas na serveru*”.
  - **Získat jméno chatbota**
    - Vrátí jméno chatbota
    - Je zavolán za předpokladu, že vstupní text obsahuje klíčové slovo “name”
    - Vrací výstup ve formátu “Chatbots name is *jméno chatbota*”
