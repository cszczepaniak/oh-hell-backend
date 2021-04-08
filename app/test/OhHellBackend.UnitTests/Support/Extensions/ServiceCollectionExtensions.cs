using Amazon.S3;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;
using OhHellBackend.UnitTests.Support.S3;

namespace OhHellBackend.DataAccess.Extensions
{
    public static class ServiceCollectionExtensions
    {
        public static IServiceCollection AddDataAccessFakes(this IServiceCollection services)
        {
            services
                .Replace(ServiceDescriptor.Singleton(typeof(IAmazonS3), p => p.GetRequiredService<FakeS3Client>()))
                .AddSingleton<FakeS3Client>();
            return services;
        }
    }
}