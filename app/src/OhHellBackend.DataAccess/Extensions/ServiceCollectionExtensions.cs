using Amazon.S3;
using Microsoft.Extensions.DependencyInjection;
using OhHellBackend.DataAccess.Games.Repositories;
using OhHellBackend.DataAccess.Services;

namespace OhHellBackend.DataAccess.Extensions
{
    public static class ServiceCollectionExtensions
    {
        public static IServiceCollection AddDataAccess(this IServiceCollection services)
        {
            services
                .AddTransient<IAmazonS3, AmazonS3Client>()
                .AddTransient<IS3Service, S3Service>()
                .AddTransient<IGameRepository, S3GameRepository>();
            return services;
        }
    }
}