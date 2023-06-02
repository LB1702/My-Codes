public static int nbJoursMoisDuneAnnee(int mois, int annee)
{
    if (mois == 2)
    {
        if (bissextile(annee))
        {
            return 29;
        }
        else
        {
            return 28;
        }
    }
    else if (mois == 4 || mois == 6 || mois == 9 || mois == 11)
    {
        return 30;
    }
    else
    {
        return 31;
    }
}



 // Comme répondu à la question, le problème était &&. Je l'ai donc remplacé par || afin que l'on puisse si le mois était égale à l'un des mois indiqués.
 // Pour être sur que le code fontionne bien, j'ai rajouté la ligne 18 pour que si le mois ne correspondait pas à l'attente demandé, il renvoyait 31 jours