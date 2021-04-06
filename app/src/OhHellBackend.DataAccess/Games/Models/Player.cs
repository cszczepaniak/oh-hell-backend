namespace OhHellBackend.DataAccess.Games.Models
{
    public class Player
    {
        public string Name { get; set; }
        public int Score { get; set; }
        public int CurrentBid { get; set; }
        public int CurrentTricks { get; set; }
    }
}