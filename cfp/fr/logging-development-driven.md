## Logging development driven talk

### Abstract

En tant que Log Evangelist j’ai la chance d’être au contact de nombreux projets. 
Et je peux affirmer sans trop me tromper que les logs sont bien présents dans **TOUS** les dév.
Mais très souvent, cela est fait à l'instinct: messages mal formatés, niveaux incohérents, ....

Car finalement très peu d'équipes ont réfléchi à ce potentiel et comment l’exploiter

Et pourtant, les logs traquent inlassablement :

* les messages et les data permettant aux dev de fiabiliser leur code
* les erreurs et les metrics permettant aux ops d’intervenir proactivement

Et avec un peu de méthode:

* les accès pour garantir sécurité et ’intégrité
* l’usage et l’activité des users pour soutenir les décisions et la roadmap produit

Je vous invite à me rejoindre et découvrir quelques pratiques éprouvées

### Keywords

`logs`, `best practices`, `architecture`

### Description

(slides à venir/sur demande)

L'enjeu de ce talk est de sensibiliser les participants au potentiel des logs.
Le talk suit le plan suivant

**Why? Pourquoi analyser et traiter les logs**

* Developpement & Quality
* Performance & Metrics =>
* Security & Compliance
* User behavior activity stream 

**How? Comment standardiser/normaliser/industrialiser le logging**

* Severity, Format and Normalisation => Json, best practices, multine line events, namming convention of most common fields and metrics
* Standardization => code the message log // inspired from JBoss Logging Framework
* Correlation basics & Tracing => adding meta and context to the log // inspired from Google Dapper paper and Zipkin
* Centralize into one place => the syslog challenge

**What? Quels informations**

* Performance KPI and metrics => all things that can be charted (response time, change state)
* Debug vs Security audit vs Forensics vs QoS, etc.
* Real user Monitoring - Focus on the behavior

*Inspiré de*

* https://www.owasp.org/index.php/Log_review_and_management 
* https://www.owasp.org/index.php/Error_Handling,_Auditing_and_Logging 
* http://docs.oracle.com/javase/6/docs/technotes/guides/logging/overview.html
* http://www.syslog.org
* http://www.slideshare.net/anton_chuvakin/pci-dss-and-logging-what-you-need-to-know-by-dr-anton-chuvakin?qid=12b7d44f-27bc-4b0e-ae2a-efeb0362229b&v=&b=&from_search=11
* https://wiki.opendaylight.org/view/BestPractices/Logging_Best_Practices 
* http://wiki.c2.com/?LoggingBestPractices 

### Bio
#### Summary
Depuis quelques années l'écosystème Tech, je suis passionné par les analytics et les stratégies data-driven. 
En tant que Log Evangelist, j’ai la chance d’être au contact de nombreux projets et entreprises.
“Ma mission”, échanger autour du potentiel des data. Et plus particulièrement celui des logs.
“Leur force”, être présent dans **TOUS** nos développements.

#### Qualifications

Speaker à diverses conférences et meetup depuis 2016. 
Organisateur du Docker Meetup Tours. 

Meetup

* Paris RB
* Paris Monitoring
* Paris NodeJS

Conferences

* Quickie / JUG Summer
* Talk / DevOps-Days Marseille
* Quickie / Dot(JS/Scale/Go)
* Talk / Paris Big Data Event

References

* https://www.youtube.com/watch?v=7Wpz_6J-49E
* https://youtu.be/509ICC4SKN0?t=3h49m20s

Intervenant sur des formations Big Data (Hadoop/Spark/Impala) (1 à 2 fois/an)
Intervenant régulier dans les entreprises pour du training aux équipes de dev autour du logs.
