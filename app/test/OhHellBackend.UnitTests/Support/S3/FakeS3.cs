using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using Amazon.S3.Model;
using OhHellBackend.DataAccess.Services;

namespace OhHellBackend.UnitTests.Support.S3
{
    public class FakeS3 : IS3Service
    {
        private readonly List<PutObjectRequest> _putObjectRequests;
        public IEnumerable<PutObjectRequest> PutObjectRequests => _putObjectRequests;
        public FakeS3()
        {
            _putObjectRequests = new List<PutObjectRequest>();
        }
        public async Task<PutObjectResponse> PutObjectAsync(
            PutObjectRequest request,
            CancellationToken _
        )
        {
            _putObjectRequests.Add(request);
            return await Task.FromResult(new PutObjectResponse());
        }
    }
}