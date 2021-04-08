using System;
using System.IO;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using OhHellBackend.DataAccess.Services;
using OhHellBackend.UnitTests.Support;
using OhHellBackend.UnitTests.Support.S3;
using Xunit;

namespace OhHellBackend.UnitTests.S3
{
    public class S3ServiceTest
    {
        private readonly IS3Service _service;
        private readonly FakeS3Client _fakeS3;
        public S3ServiceTest()
        {
            var provider = ServiceProviderFactory.Create();
            _fakeS3 = provider.GetRequiredService<FakeS3Client>();
            _service = provider.GetRequiredService<IS3Service>();
        }

        [Fact]
        public async Task CreateJsonAsync_CreatesPutRequest()
        {
            var obj = new { Something = "data" };
            _fakeS3.ConfigureSuccessfulPutRequest("bucket");

            await _service.UploadJsonAsync("bucket", "key", obj);

            var request = Assert.Single(_fakeS3.PutObjectRequests);
            Assert.Equal("application/json", request.ContentType);
            Assert.Equal("bucket", request.BucketName);
            Assert.Equal("key.json", request.Key);
            var json = await new StreamReader(request.InputStream).ReadToEndAsync();
            Assert.Equal("{\"Something\":\"data\"}", json);
        }

        [Fact]
        public async Task CreateJsonAsyc_ThrowsWhenS3ReturnsErrorCode()
        {
            _fakeS3.ConfigureFailingPutRequest("bucket");

            await Assert.ThrowsAsync<HttpRequestException>(async () => await _service.UploadJsonAsync("bucket", "key", null));
        }
    }
}