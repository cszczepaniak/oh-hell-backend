using System.Collections.Generic;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace OhHellBackend.DataAccess.Games.Models
{
    public class Game
    {
        public string Id { get; set; }
        public IEnumerable<Player> Players { get; set; }
        public int Round { get; set; }
        public int NumberOfCards { get; set; }
        public int DealerIndex { get; set; }
        [JsonConverter(typeof(StringEnumConverter))]
        public GameSettings Settings { get; set; }

    }
}