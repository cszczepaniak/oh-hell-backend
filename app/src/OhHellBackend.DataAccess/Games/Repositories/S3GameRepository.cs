using System.Threading.Tasks;
using OhHellBackend.DataAccess.Games.Models;

namespace OhHellBackend.DataAccess.Games.Repositories
{
    public class S3GameRepository : IGameRepository
    {
        public Task<string> CreateAsync(Game game)
        {
            throw new System.NotImplementedException();
        }
    }
}