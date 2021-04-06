using System;
using System.Threading;
using System.Threading.Tasks;
using Amazon.S3;
using Amazon.S3.Model;

namespace OhHellBackend.DataAccess.Services
{
    public class S3Service : IS3Service
    {
        public Task<PutObjectResponse> PutObjectAsync(PutObjectRequest request, CancellationToken cancellationToken = default)
        {
            throw new NotImplementedException();
        }
    }
}