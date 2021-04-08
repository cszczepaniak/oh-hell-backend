using System;
using System.Threading.Tasks;

namespace OhHellBackend.DataAccess.Services
{
    public interface IS3Service
    {
        Task UploadJsonAsync(string bucketName, string key, object obj);
    }
}