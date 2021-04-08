using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;
using OhHellBackend.DataAccess.Services;
using OhHellBackend.UnitTests.Support.S3;

namespace OhHellBackend.DataAccess.Extensions
{
    public static class ServiceCollectionExtensions
    {
        public static IServiceCollection AddDataAccessFakes(this IServiceCollection services)
        {
            services.Replace(new ServiceDescriptor(typeof(IS3Service), typeof(FakeS3)));
            return services;
        }
    }
}