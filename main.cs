static void Main(string[] args)
{
    Console.Write("mois=");
    int mois;
    if (!int.TryParse(Console.ReadLine(), out mois) || mois < 1 || mois > 12)
    {
        Console.WriteLine("Mois non valide. Entrer un mois valide compris entre 1 et 12.");
        Console.ReadLine();
        return;
    }

    Console.Write("année=");
    int annee;
    if (!int.TryParse(Console.ReadLine(), out annee) || annee < 1)
    {
        Console.WriteLine("Année non validée. Entrer une année valide.");
        Console.ReadLine();
        return;
    }

    Console.WriteLine("nombres de jours = " + nbJoursMoisDuneAnnee(mois, annee));
    Console.ReadLine();
}



// Comme vous pouvez constater, à la ligne 5 j'ai modifié int.Parse par int.TryParse car cela permet au code de convertir les entrées de l'utilisateur en entier. Si les valeurs sont
// supérieur à 12 ou bien inférieur à 1 ou bien si la conversion se passe mal, il y aura un message qui s'affichera afin de comprendre le problème et le problème s'arretera.