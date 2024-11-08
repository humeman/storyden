package semdexer

import (
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"go.uber.org/fx"

	"github.com/Southclaws/storyden/app/resources/datagraph/semdex"
	"github.com/Southclaws/storyden/app/services/semdex/semdexer/refhydrate"
	"github.com/Southclaws/storyden/app/services/semdex/semdexer/simplesearch"
	"github.com/Southclaws/storyden/app/services/semdex/semdexer/weaviate_semdexer"
	"github.com/Southclaws/storyden/internal/config"
	weaviate_infra "github.com/Southclaws/storyden/internal/infrastructure/weaviate"
)

func newSemdexer(
	cfg config.Config,
	wc *weaviate.Client,
	simpleSearcher *simplesearch.ParallelSearcher,

	weaviateClassName weaviate_infra.WeaviateClassName,
	hydrator *refhydrate.Hydrator,
) semdex.Semdexer {
	if !cfg.SemdexEnabled {
		return &semdex.OnlySearcher{Searcher: simpleSearcher}
	}

	return weaviate_semdexer.New(wc, weaviateClassName, hydrator)
}

func Build() fx.Option {
	return fx.Provide(
		refhydrate.New,
		simplesearch.NewParallelSearcher,
		fx.Annotate(
			newSemdexer,
			fx.As(new(semdex.Semdexer)),
			fx.As(new(semdex.Indexer)),
			fx.As(new(semdex.Searcher)),
			fx.As(new(semdex.Recommender)),
			fx.As(new(semdex.Retriever)),
			fx.As(new(semdex.RelevanceScorer)),
			fx.As(new(semdex.Summariser)),
		),
	)
}
