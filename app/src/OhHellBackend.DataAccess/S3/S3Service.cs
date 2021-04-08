using System;
using System.IO;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Amazon.S3;
using Amazon.S3.Model;
using Newtonsoft.Json;

namespace OhHellBackend.DataAccess.Services
{
    public class S3Service : IS3Service
    {
        private readonly IAmazonS3 _s3Client;

        public S3Service(IAmazonS3 s3Client)
        {
            _s3Client = s3Client;
        }

        public async Task UploadJsonAsync(string bucketName, string key, object obj)
        {
            var json = JsonConvert.SerializeObject(obj);
            var bytes = Encoding.ASCII.GetBytes(json);
            var request = new PutObjectRequest
            {
                BucketName = bucketName,
                Key = $"{key}.json",
                ContentType = "application/json",
                InputStream = new MemoryStream(bytes),
            };
            var result = await _s3Client.PutObjectAsync(request);
            new HttpResponseMessage(result.HttpStatusCode).EnsureSuccessStatusCode();
        }
    }
}