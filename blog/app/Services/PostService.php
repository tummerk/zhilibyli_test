<?php

namespace App\Services;

use App\Contracts\Services\PostServiceInterface;
use App\Http\Resources\PostResource;
use App\Models\Post;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Pagination\LengthAwarePaginator;
use Illuminate\Support\Str;

class PostService implements PostServiceInterface//инкапсулирует всю работу с бд,
                  // можно было бы конечно на таком простом проекте сделать и в самом контроллере
{
    public function getAllPosts(): LengthAwarePaginator
    {
        return Post::Paginate(15);
    }

    public function createPost(array $data): Post
    {
        return Post::create($data);
    }

    public function updatePost(string $id, array $data): Post
    {
        $post=Post::findOrFail($id); //для обработки ситуации когда пост не найден
        $post->update($data); // с RMB была бы одна операция, иду на это ради кастомных ответов и слоистой архитектуры

        return $post;
    }

    public function deletePost(string $id): void
    {
        $post=Post::findOrFail($id);
        $post->delete();
    }

    public function getPostsCount(): int
    {
        return Post::count();
    }

    public function getPostById(string $id): Post
    {
        return Post::findOrFail($id);
    }
}
