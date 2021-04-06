using System.Threading;
using System.Threading.Tasks;
using Amazon.S3.Model;

namespace OhHellBackend.DataAccess.Services
{
    public interface IS3Service
    {
        Task<PutObjectResponse> PutObjectAsync(PutObjectRequest request, CancellationToken cancellationToken);
    }
}