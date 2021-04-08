using System;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using OhHellBackend.DataAccess.Games.Models;
using OhHellBackend.DataAccess.Games.Repositories;
using OhHellBackend.UnitTests.Support;
using OhHellBackend.UnitTests.Support.S3;
using Xunit;

namespace OhHellBackend.UnitTests
{
    public class S3GameRepositoryTest
    {
        private readonly FakeS3Client _s3;
        private readonly IGameRepository _repository;
        public S3GameRepositoryTest()
        {
            var provider = ServiceProviderFactory.Create();
            _s3 = provider.GetRequiredService<FakeS3Client>();
            _repository = provider.GetRequiredService<IGameRepository>();
        }
        [Fact]
        public async Task CreateAsync_ShouldIssuePutRequestToS3()
        {
        }
    }
}
