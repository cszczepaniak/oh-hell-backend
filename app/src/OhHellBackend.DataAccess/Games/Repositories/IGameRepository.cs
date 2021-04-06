using System.Threading.Tasks;
using OhHellBackend.DataAccess.Games.Models;

namespace OhHellBackend.DataAccess.Games.Repositories
{
    public interface IGameRepository
    {
        Task<string> CreateAsync(Game game);
    }
}